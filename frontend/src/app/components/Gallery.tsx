import Link from 'next/link';
import {Card, CardBody, Masonry, ResponsiveMasonry} from '../common'
import { getAllQuotes } from '../services/api';
import { ListQuotesResponse } from '@/generated/buf/kotobako/v1/kotobako_pb';

const Gallery = async () => {
    const quotesResponse: ListQuotesResponse = await getAllQuotes();
    const quotes = quotesResponse.quotes;

    const childElements = quotes.map((quote) => (
        <Card key={quote.quoteId}>
            <Link href={`/detail/${quote.quoteId}`}>
                <CardBody padding={Math.random()*(60-20)+20}>
                    {quote.sentence}
                </CardBody>
            </Link>
        </Card>
    ))
    return (
        <main className="main">
            <ResponsiveMasonry columnsCountBreakPoints={{ 350: 1, 700: 2, 900: 3, 1100: 4}}>
                <Masonry gutter='20px'>
                {childElements}
                </Masonry>
            </ResponsiveMasonry>
        </main>
    )
}

export default Gallery;