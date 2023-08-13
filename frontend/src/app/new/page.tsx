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
<<<<<<< HEAD
import { postQuote } from "../services/api";
import { PostQuoteRequest } from "@/generated/buf/kotobako/v1/kotobako_pb";
=======
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64

export default function CreateArticle() {
    const router = useRouter();
    const [isPending, startTransition] = useTransition();
    const [quote, setTtile] = useState("");
    const [quote_source, setContent] = useState("");
    const [quote_source_name, setQuoteSourceName] = useState("book");
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
<<<<<<< HEAD

        e.preventDefault();
        setLoading(true);
        await postQuote(new PostQuoteRequest({ sentence: quote, quoteSourceName: quote_source, quoteMediaType: quote_source_name, authorName: "hoge" }));
=======
        e.preventDefault();
        setLoading(true);
        // await fetch("/v1/quotes", {
        //     method: "POST",
        //     headers: {
        //         "Content-Type": "application/json",
        //     },
        //     body: JSON.stringify({ title, content }),
        // });
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64
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
<<<<<<< HEAD
                            <Radio value="book">小説・漫画</Radio>
=======
                            <Radio value="book">本</Radio>
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64
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