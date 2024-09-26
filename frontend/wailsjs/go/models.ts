export namespace main {
	
	export class ProjectDetails {
	
	
	    static createFrom(source: any = {}) {
	        return new ProjectDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

