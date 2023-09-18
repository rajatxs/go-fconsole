export namespace models {
	
	export class PostCoverImage {
	    id: string;
	    path: string;
	    refName: string;
	    refUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new PostCoverImage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.refName = source["refName"];
	        this.refUrl = source["refUrl"];
	    }
	}
	export class PostMetadataDocument {
	    _id: number[];
	    title: string;
	    slug: string;
	    desc: string;
	    tags: string[];
	    topic: string;
	    stars: number;
	    authorId: number[];
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    coverImage?: PostCoverImage;
	
	    static createFrom(source: any = {}) {
	        return new PostMetadataDocument(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this._id = source["_id"];
	        this.title = source["title"];
	        this.slug = source["slug"];
	        this.desc = source["desc"];
	        this.tags = source["tags"];
	        this.topic = source["topic"];
	        this.stars = source["stars"];
	        this.authorId = source["authorId"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.coverImage = this.convertValues(source["coverImage"], PostCoverImage);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

