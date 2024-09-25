export namespace application {
	
	export class MessagesStruct {
	    SkillExistsMessage: string;
	    SkillAddedMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new MessagesStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SkillExistsMessage = source["SkillExistsMessage"];
	        this.SkillAddedMessage = source["SkillAddedMessage"];
	    }
	}
	export class Skill {
	    Name: string;
	    SVG: string;
	
	    static createFrom(source: any = {}) {
	        return new Skill(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.SVG = source["SVG"];
	    }
	}

}

