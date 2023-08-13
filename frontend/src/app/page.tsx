"use client"
import useSWR from "swr"
import Gallery from './components/Gallery';

export default function Main() {
    return (
        <main className="main">
            <Gallery/>
        </main>
    )
}