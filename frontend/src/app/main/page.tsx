"use client";
import { getJSON } from '../services/api'
import useSWR from "swr"
import Gallery from '../components/Gallery';

export default function Main() {
    return (
        <main className="main">
            this is main
            <main className="styles.main">
                <Gallery/>
            </main>
        </main>
    )
}