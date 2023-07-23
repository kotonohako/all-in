import { Card, CardBody, CardHeader, Heading, Text } from "@/app/common";
import { Quote } from "@/app/types";

export default function QuoteDetail({quote}:{quote:Quote}) {
    return (
        <Card as="article">
            <CardHeader>
                <Heading as="h1">{quote.sentence}</Heading>
            </CardHeader>
            <CardBody>
                <Text as="p" fontSize="md">
                    {quote.quote_source_name}
                </Text>
            </CardBody>
        </Card>
    )
}