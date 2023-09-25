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

export namespace types {
	
	export class AppPublicConfigVariables {
	    ENV: string;
	    ADMIN_ID: string;
	    CLOUDINARY_ID: string;
	
	    static createFrom(source: any = {}) {
	        return new AppPublicConfigVariables(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ENV = source["ENV"];
	        this.ADMIN_ID = source["ADMIN_ID"];
	        this.CLOUDINARY_ID = source["CLOUDINARY_ID"];
	    }
	}
	export class AppVersions {
	    app: string;
	    date: string;
	    wails: string;
	    go: string;
	    webview2: string;
	    os: string;
	    arch: string;
	    uname: string;
	    homedir: string;
	
	    static createFrom(source: any = {}) {
	        return new AppVersions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.app = source["app"];
	        this.date = source["date"];
	        this.wails = source["wails"];
	        this.go = source["go"];
	        this.webview2 = source["webview2"];
	        this.os = source["os"];
	        this.arch = source["arch"];
	        this.uname = source["uname"];
	        this.homedir = source["homedir"];
	    }
	}
	export class CreatePostPayload {
	    title: string;
	    slug: string;
	    desc: string;
	    tags: string[];
	    topic: string;
	    body: string;
	    public: boolean;
	    coverImageId: string;
	    coverImagePath: string;
	    coverImageRefName: string;
	    coverImageRefUrl: string;
	    authorId: string;
	
	    static createFrom(source: any = {}) {
	        return new CreatePostPayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.slug = source["slug"];
	        this.desc = source["desc"];
	        this.tags = source["tags"];
	        this.topic = source["topic"];
	        this.body = source["body"];
	        this.public = source["public"];
	        this.coverImageId = source["coverImageId"];
	        this.coverImagePath = source["coverImagePath"];
	        this.coverImageRefName = source["coverImageRefName"];
	        this.coverImageRefUrl = source["coverImageRefUrl"];
	        this.authorId = source["authorId"];
	    }
	}
	export class GetPostsMetadataOptions {
	    private: boolean;
	    topic: string;
	    sortBy: string;
	    limit: number;
	    skip: number;
	
	    static createFrom(source: any = {}) {
	        return new GetPostsMetadataOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.private = source["private"];
	        this.topic = source["topic"];
	        this.sortBy = source["sortBy"];
	        this.limit = source["limit"];
	        this.skip = source["skip"];
	    }
	}

}

