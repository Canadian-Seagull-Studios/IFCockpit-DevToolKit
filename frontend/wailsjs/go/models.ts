export namespace runtime {
	
	export class EnvironmentInfo {
	    buildType: string;
	    platform: string;
	    arch: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvironmentInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buildType = source["buildType"];
	        this.platform = source["platform"];
	        this.arch = source["arch"];
	    }
	}

}

