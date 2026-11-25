export namespace core {
	
	export class ApplyResult {
	    Locale: string;
	    LocaleName: string;
	    BackupPath: string;
	    CleanBackup: string;
	    WasRunning: boolean;
	    Forced: boolean;
	    Relaunched: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ApplyResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Locale = source["Locale"];
	        this.LocaleName = source["LocaleName"];
	        this.BackupPath = source["BackupPath"];
	        this.CleanBackup = source["CleanBackup"];
	        this.WasRunning = source["WasRunning"];
	        this.Forced = source["Forced"];
	        this.Relaunched = source["Relaunched"];
	    }
	}
	export class LocaleInfo {
	    code: string;
	    name: string;
	    englishName: string;
	    translations: number;
	    coverage: number;
	
	    static createFrom(source: any = {}) {
	        return new LocaleInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.name = source["name"];
	        this.englishName = source["englishName"];
	        this.translations = source["translations"];
	        this.coverage = source["coverage"];
	    }
	}
	export class RestoreResult {
	    AlreadyEnglish: boolean;
	    BackupPath: string;
	    WasRunning: boolean;
	    Forced: boolean;
	    Relaunched: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RestoreResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AlreadyEnglish = source["AlreadyEnglish"];
	        this.BackupPath = source["BackupPath"];
	        this.WasRunning = source["WasRunning"];
	        this.Forced = source["Forced"];
	        this.Relaunched = source["Relaunched"];
	    }
	}
	export class Status {
	    AppDir: string;
	    ResourcesDir: string;
	    AsarPath: string;
	    CleanBackupPath: string;
	    StatePath: string;
	    AppVersion: string;
	    CurrentLanguage: string;
	    LocaleName: string;
	    SessionMetrics: boolean;
	    EffortSlider: boolean;
	    HasCleanBackup: boolean;
	    CleanBackupStatus: string;
	    LastPatchedAt?: string;
	    StateIsStale: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Status(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AppDir = source["AppDir"];
	        this.ResourcesDir = source["ResourcesDir"];
	        this.AsarPath = source["AsarPath"];
	        this.CleanBackupPath = source["CleanBackupPath"];
	        this.StatePath = source["StatePath"];
	        this.AppVersion = source["AppVersion"];
	        this.CurrentLanguage = source["CurrentLanguage"];
	        this.LocaleName = source["LocaleName"];
	        this.SessionMetrics = source["SessionMetrics"];
	        this.EffortSlider = source["EffortSlider"];
	        this.HasCleanBackup = source["HasCleanBackup"];
	        this.CleanBackupStatus = source["CleanBackupStatus"];
	        this.LastPatchedAt = source["LastPatchedAt"];
	        this.StateIsStale = source["StateIsStale"];
	    }
	}

}

export namespace main {

	export class UpdateCheckResult {
	    release?: selfupdate.Release;
	    error?: string;
	    retryAt?: number;

	    static createFrom(source: any = {}) {
	        return new UpdateCheckResult(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.release = this.convertValues(source["release"], selfupdate.Release);
	        this.error = source["error"];
	        this.retryAt = source["retryAt"];
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace selfupdate {

	export class Release {
	    version: string;
	    notes: string;
	    available: boolean;
	    retryAt?: number;

	    static createFrom(source: any = {}) {
	        return new Release(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.notes = source["notes"];
	        this.available = source["available"];
	        this.retryAt = source["retryAt"];
	    }
	}

}

