"use client";

import { useState, useTransition } from "react";
import { useRouter } from "next/navigation";
import {
    Heading,
    FormControl,
    FormLabel,
    Input,
    Textarea,
    Button,
    Radio,
    RadioGroup,
    Stack
} from "../common";

export default function CreateArticle() {
    const router = useRouter();
    const [isPending, startTransition] = useTransition();
    const [quote, setTtile] = useState("");
    const [quote_source, setContent] = useState("");
    const [quote_source_name, setQuoteSourceName] = useState("book");
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        // await fetch("/v1/quotes", {
        //     method: "POST",
        //     headers: {
        //         "Content-Type": "application/json",
        //     },
        //     body: JSON.stringify({ title, content }),
        // });
        setLoading(false);
        router.push("/");
        startTransition(() => {
            router.refresh;
        })
    }
    return (
        <div>
            <Heading mb={4}>投稿する</Heading>
            <form onSubmit={handleSubmit}>
                <FormControl>
                    <FormLabel>コトバ</FormLabel>
                    <Textarea value={quote} onChange={(e) => setTtile(e.target.value)} />
                    
                    <FormLabel>引用元</FormLabel>
                    <Input value={quote_source} onChange={(e) => setContent(e.target.value)} />

                    <RadioGroup onChange={setQuoteSourceName} value={quote_source_name} defaultValue="book"> 
                        <Stack direction='row'>
                            <Radio value="book">本</Radio>
                            <Radio value="movie">映画</Radio>
                            <Radio value="game">ゲーム</Radio>
                            <Radio value="music">歌詞</Radio>
                            <Radio value="copywrite">広告</Radio>
                            <Radio value="etc">その他</Radio>
                        </Stack>
                    </RadioGroup>
                    <Button
                        type="submit"
                        color="white"
                        bg="orange.400"
                        isLoading={loading || isPending}
                        mt={4}
                    >
                        作成
                    </Button>
                </FormControl>
            </form>
        </div>
    )
}