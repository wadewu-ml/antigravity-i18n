const assert = require('node:assert/strict');
const { mkdirSync, readFileSync } = require('node:fs');
const { join } = require('node:path');
const { DatabaseSync } = require('node:sqlite');

function varint(value) {
    const bytes = [];
    while (value >= 128) {
        bytes.push(value % 128 + 128);
        value = Math.floor(value / 128);
    }
    return Buffer.from([...bytes, value]);
}

function scalar(field, value) {
    return Buffer.concat([varint(field * 8), varint(value)]);
}

function message(field, ...values) {
    const body = Buffer.concat(values);
    return Buffer.concat([varint(field * 8 + 2), varint(body.length), body]);
}

function generation(output) {
    return message(1,
        message(4, scalar(2, 200), scalar(3, output), scalar(5, 800)),
        message(9, message(10, scalar(4, 10000))),
        message(11, scalar(1, 1), scalar(2, 250000000)),
        message(12, scalar(1, 5)));
}

async function test() {
    const home = process.argv[2];
    const directory = join(home, '.gemini', 'antigravity', 'conversations');
    mkdirSync(directory, { recursive: true });
    const id = '10000000-0000-0000-0000-000000000000';
    const db = new DatabaseSync(join(directory, id + '.db'));
    try {
        db.exec('PRAGMA journal_mode = WAL; CREATE TABLE gen_metadata (idx INTEGER PRIMARY KEY, data BLOB)');
        const insert = db.prepare('INSERT INTO gen_metadata VALUES (?, ?)');
        insert.run(1, generation(1200));
        insert.run(2, Buffer.from([10, 5, 1]));
        let handle;
        const install = new Function('require', readFileSync('../../assets/patches/metrics-main.jsfrag', 'utf8'));
        install((module) => {
            if (module === './utils') return { isLocalHubTopFrame: (event) => event.local };
            if (module === 'electron') return {
                app: { getPath: () => home },
                ipcMain: { handle: (_, fn) => { handle = fn; } },
            };
            return require(module);
        });
        const event = {local: true};
        assert.deepEqual(await handle(event, id), {
            used: 1000, limit: 10000, speed: 240, cacheRate: 0.8, firstToken: 1.25,
        });
        assert.equal(await handle({local: false}, id), null);
        assert.equal(await handle(event, '../unknown'), null);
        assert.equal(await handle(event, '20000000-0000-0000-0000-000000000000'), null);
        insert.run(3, generation(1400));
        const updated = await handle(event, id);
        assert.equal(updated.speed, 280);
        assert.deepEqual(await handle(event, id), updated);
    } finally {
        db.close();
    }
}

test().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
