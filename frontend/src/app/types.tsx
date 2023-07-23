export type Quote = {
    id: number,
    quote_media_type: 'book'|'movie'|'game'|'music'|'copywrite'|'etc',
    quote_source_name: string,
    sentence: string,
    speaker_name?: string;
};