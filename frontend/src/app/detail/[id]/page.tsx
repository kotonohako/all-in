"use client";
import Gallery from '../../components/Gallery';

// type PageProps = {
//     params: {
//         id: number;
//     }
// }

const Page = ({ params }: { params: { id: string } }) => {
    return (
        <h1>
            this is detail{params.id}
            <main className="main">
                <Gallery/>
            </main>
        </h1>
    )
}

export default Page;