"use client";
import Link from 'next/link';
import Masonry from 'react-responsive-masonry'

const Gallery = () => {
    const nums = [...Array(10)].map((_, i) => (i));
    const childElements = nums.map(
        function (element) {
            const sentenceDOM = (
                <p key={`sentence-${element}`}>
                    {`${element + 1} elements sentences\n`.repeat(element + 1)}
                </p>
            )
            return (
                <Link href="/detail/hogehoge" key={element}>
                    <div className="card">
                        {sentenceDOM}
                    </div>
                </Link>
            )
        }
    )
    return (
        <main className="main">
            <Masonry>
                {childElements}
            </Masonry>
        </main>
    )
}

export default Gallery;