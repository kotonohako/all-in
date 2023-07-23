import Link from 'next/link';
import {Card, CardBody, Masonry, ResponsiveMasonry} from '../common'
import { Quote } from '../types';

async function getAllQuotes() {
    const res = await fetch("http://localhost:8080/v1/quotes")
    if (!res.ok) {
        throw new Error("failed to fetch quotes");
    }
    const data = await res.json();
    return data as Quote[];
}

const Gallery = async () => {
    const quotes = await getAllQuotes();
    const childElements = quotes.map((quote) => (
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
            <ResponsiveMasonry columnsCountBreakPoints={{ 350: 1, 700: 2, 900: 3, 1100: 4}}>
                <Masonry gutter='20px'>
                {childElements}
                </Masonry>
            </ResponsiveMasonry>
        </main>
    )
}

export default Gallery;