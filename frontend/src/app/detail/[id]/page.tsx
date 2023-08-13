"use client";
<<<<<<< HEAD
import Gallery from '../../components/Gallery';
import QuoteDetail from './QuoteDetail';
import {Heading, Spacer, VStack} from '../../common'
import { GetQuoteResponse, Quote } from '@/generated/buf/kotobako/v1/kotobako_pb';
import { getQuote } from '@/app/services/api';

export default async function Page({ params }: { params: { id: string } }) {
    const quotesResponse: GetQuoteResponse = await getQuote(params.id);
    const quote = quotesResponse as Quote;
=======
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
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64
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