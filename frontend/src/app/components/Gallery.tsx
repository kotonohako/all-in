import Link from 'next/link';
import Masonry from 'react-responsive-masonry'
import {Card, CardBody} from '@chakra-ui/react'

type Quote = {
    id: number,
    quote_media_type: string,
    quote_source_name: string,
    sentence: string,
    speaker_name?: string;
};

const Gallery = async () => {
    const data = await fetch("http://localhost:8080/v1/quotes")
    const quoteInfo: Quote[] = await data.json()
    const childElements = quoteInfo.map((quote) => (
        <Card key={quote.id}>
            <CardBody>
            <Link href={`/detail/${quote.id}`}>
                {quote.sentence}
            </Link>
            </CardBody>
        </Card>
    ))
    return (
        <main className="main">
            <Masonry gutter='20px'>
                {childElements}
            </Masonry>
        </main>
    )
}

export default Gallery;