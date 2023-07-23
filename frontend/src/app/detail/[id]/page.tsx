"use client";
import { Quote } from '@/app/types';
import Gallery from '../../components/Gallery';

async function getQuote(id:string) {
    const res = await fetch(`http://localhost:8080/v1/quotes/${id}`)
    if (!res.ok) {
        throw new Error("failed to fetch quotes");
    }
    const data = await res.json();
    return data as Quote;
}

export default async function Page({ params }: { params: { id: string } }){
    const quote = await getQuote(params.id)
    return (
        <h1>
            {quote.sentence}
            {quote.quote_source_name}
            {quote.quote_media_type}
            <main className="main">
                <Gallery/>
            </main>
        </h1>
    )
}