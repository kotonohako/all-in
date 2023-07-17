"use client";
import Gallery, { Quote } from '../../components/Gallery';

// type PageProps = {
//     params: {
//         id: number;
//     }
// }

const Page = async ({ params }: { params: { id: string } }) => {
    const data = await fetch(`http://localhost:8080/v1/quotes/${params.id}`)
    const quote: Quote = await data.json()
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

export default Page;