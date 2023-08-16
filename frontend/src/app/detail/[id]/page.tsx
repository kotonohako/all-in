"use client";
import Gallery from '@/app/components/Gallery';
import QuoteDetail from '@/app/detail/[id]/QuoteDetail';
import {Heading, Spacer, VStack} from '@/app/common'
import { GetQuoteResponse, Quote } from '@/generated/buf/kotobako/v1/kotobako_pb';
import { getQuote } from '@/app/services/api';

export default async function Page({ params }: { params: { id: string } }) {
    const quotesResponse: GetQuoteResponse = await getQuote(params.id);
    const quote = quotesResponse as Quote;
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