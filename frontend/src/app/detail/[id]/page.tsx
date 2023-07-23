"use client";
import { Quote } from '@/app/types';
import Gallery from '../../components/Gallery';
import QuoteDetail from './QuoteDetail';
import {Heading, Spacer, VStack} from '../../common'

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
        <div>
            <QuoteDetail quote={quote} />
            <Heading as="h2" mt={8} mb={4} fontSize={24}>
                その他
            </Heading>
            <Gallery/>
        </div>
    )
}