export namespace main {
	
	export class GopherItem {
	    type: string;
	    display: string;
	    selector: string;
	    host: string;
	    port: string;
	
	    static createFrom(source: any = {}) {
	        return new GopherItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.display = source["display"];
	        this.selector = source["selector"];
	        this.host = source["host"];
	        this.port = source["port"];
	    }
	}
	export class GopherResponse {
	    items: GopherItem[];
	    raw: string;
	    err: string;
	    contentType: string;
	
	    static createFrom(source: any = {}) {
	        return new GopherResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], GopherItem);
	        this.raw = source["raw"];
	        this.err = source["err"];
	        this.contentType = source["contentType"];
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

