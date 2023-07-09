"use client";
import styles from './page.module.css'
import Link from 'next/link'
import GalleryComponent from '../gallery/GalleryComponent';

export default function Main() {
    return (
        <main className={styles.main}>
            <div>gallery</div>
            <Link href="/gallery">Go to gallery</Link>
            <GalleryComponent sentence="fuga"/>
        </main>
    )
}