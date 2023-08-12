"use client";
import styles from './page.module.css'
import { getJSON, hello } from '../services/api'
import useSWR from "swr"
import { KotobakoService } from '@/generated/buf/kotobako/v1/kotobako_connect';

export default function About() {
    const { data, error } = useSWR([KotobakoService.methods.health.name], hello)
    return (
        <main className={styles.main}>
            <div>{JSON.stringify(data)}</div>
        </main>
    )
}