import Link from 'next/link';
import Masonry from 'react-responsive-masonry'

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
        <div key={quote.id} className='card'>
            <Link href={`/detail/${quote.id}`}>
                {quote.sentence}
            </Link>
        </div>
    ))
    return (
        <main className="main">
            <Masonry>
                {childElements}
            </Masonry>
        </main>
    )
}

export default Gallery;