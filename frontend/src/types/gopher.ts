export interface GopherItem {
    type: string;
    display: string;
    selector: string;
    host: string;
    port: string;
}

export interface GopherResponse {
    items: GopherItem[];
    raw: string;
    err?: string;
    contentType?: string;
}
