"use client";
import styles from './page.module.css'
import Link from 'next/link'
import Gallery from '../gallery/page';

export default function Main() {
    return (
        <main className={styles.main}>
            <div>gallery</div>
            <Link href="/gallery">Go to gallery</Link>
            <Gallery/>
        </main>
    )
}