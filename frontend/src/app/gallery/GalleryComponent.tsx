"use client";
import styles from './page.module.css'
import Masonry from 'react-responsive-masonry'

type Props = {
    sentence : string
}

export default function GalleryComponent(props:Props) {
    const nums = [...Array(10)].map((_, i) => (i));
    const childElements = nums.map(
        function (element) {
            const sentenceDOM = (
                <p key={`sentence-${element}`}>
                    {`${element + 1} ${props.sentence}\n`.repeat(element + 1)}
                </p>
            )
            return (
                <div className={styles.card} key={element}>
                    {sentenceDOM}
                </div>
            )
        }
    );
    return (
        <main className={styles.main}>
            <Masonry>
                {childElements}
            </Masonry>
        </main>
    )
}