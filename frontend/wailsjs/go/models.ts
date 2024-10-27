export namespace main {
	
	export enum Event {
	    STRING_CHANGED = "stringChanged",
	}

}

export namespace scoreboard {
	
	export class Player {
	    Name: string;
	    TeamName: string;
	    Score: number;
	
	    static createFrom(source: any = {}) {
	        return new Player(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.TeamName = source["TeamName"];
	        this.Score = source["Score"];
	    }
	}
	export class Scoreboard {
	    BestOf: number;
	    RoundName: string;
	    Players: Player[];
	
	    static createFrom(source: any = {}) {
	        return new Scoreboard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.BestOf = source["BestOf"];
	        this.RoundName = source["RoundName"];
	        this.Players = this.convertValues(source["Players"], Player);
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

