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
	    format: string;
	    authorId: number[];
	    license: string;
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
	        this.format = source["format"];
	        this.authorId = source["authorId"];
	        this.license = source["license"];
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
	export class PostRelatedView {
	    _id: number[];
	    title: string;
	    slug: string;
	    desc: string;
	    format: string;
	    stars: number;
	    public: boolean;
	    coverImage?: PostCoverImage;
	    authorId: number[];
	    license: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new PostRelatedView(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this._id = source["_id"];
	        this.title = source["title"];
	        this.slug = source["slug"];
	        this.desc = source["desc"];
	        this.format = source["format"];
	        this.stars = source["stars"];
	        this.public = source["public"];
	        this.coverImage = this.convertValues(source["coverImage"], PostCoverImage);
	        this.authorId = source["authorId"];
	        this.license = source["license"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
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
	export class PostObjectView {
	    _id: number[];
	    title: string;
	    slug: string;
	    desc: string;
	    tags: string[];
	    topic: string;
	    body: {[key: string]: any};
	    format: string;
	    stars: number;
	    public: boolean;
	    coverImage?: PostCoverImage;
	    authorId: number[];
	    license: string;
	    relatedPosts: PostRelatedView[];
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new PostObjectView(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this._id = source["_id"];
	        this.title = source["title"];
	        this.slug = source["slug"];
	        this.desc = source["desc"];
	        this.tags = source["tags"];
	        this.topic = source["topic"];
	        this.body = source["body"];
	        this.format = source["format"];
	        this.stars = source["stars"];
	        this.public = source["public"];
	        this.coverImage = this.convertValues(source["coverImage"], PostCoverImage);
	        this.authorId = source["authorId"];
	        this.license = source["license"];
	        this.relatedPosts = this.convertValues(source["relatedPosts"], PostRelatedView);
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
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
	    body: {[key: string]: any};
	    format: string;
	    public: boolean;
	    coverImageId: string;
	    coverImagePath: string;
	    coverImageRefName: string;
	    coverImageRefUrl: string;
	    authorId: string;
	    license: string;
	    relatedPosts: string[];
	
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
	        this.format = source["format"];
	        this.public = source["public"];
	        this.coverImageId = source["coverImageId"];
	        this.coverImagePath = source["coverImagePath"];
	        this.coverImageRefName = source["coverImageRefName"];
	        this.coverImageRefUrl = source["coverImageRefUrl"];
	        this.authorId = source["authorId"];
	        this.license = source["license"];
	        this.relatedPosts = source["relatedPosts"];
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
	export class Topic {
	    name: string;
	    thumbId: string;
	    thumbPath: string;
	    public: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Topic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.thumbId = source["thumbId"];
	        this.thumbPath = source["thumbPath"];
	        this.public = source["public"];
	    }
	}
	export class UpdatePostPayload {
	    title: string;
	    slug: string;
	    desc: string;
	    tags: string[];
	    topic: string;
	    body: {[key: string]: any};
	    public: boolean;
	    coverImageId: string;
	    coverImagePath: string;
	    coverImageRefName: string;
	    coverImageRefUrl: string;
	    license: string;
	    relatedPosts: string[];
	
	    static createFrom(source: any = {}) {
	        return new UpdatePostPayload(source);
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
	        this.license = source["license"];
	        this.relatedPosts = source["relatedPosts"];
	    }
	}
	export class UploadedImageFile {
	    publicId: string;
	    assetId: string;
	    format: string;
	
	    static createFrom(source: any = {}) {
	        return new UploadedImageFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.publicId = source["publicId"];
	        this.assetId = source["assetId"];
	        this.format = source["format"];
	    }
	}

}

