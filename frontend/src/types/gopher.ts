export interface GopherItem {
    type: string;
    display: string;
    selector: string;
    host: string;
    port: string;
    description: string;
}

export interface GopherResponse {
    items: GopherItem[];
    raw: string;
    err?: string;
}

export interface ParsedGopherUrl {
    host: string;
    port: string;
    selector: string;
}