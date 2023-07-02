"use client";
import styles from './page.module.css'
import { getJSON } from '../services/api'
import useSWR from "swr"

export default function About() {
    const { data, error } = useSWR("http://localhost:8080/hello", getJSON)
    return (
        <main className={styles.main}>
            <div>{JSON.stringify(data)}</div>
        </main>
    )
}