import { Card, CardBody, CardHeader, Heading, Text } from "@/app/common";
import { Quote } from "@/generated/buf/kotobako/v1/kotobako_pb";

export default function QuoteDetail({quote}:{quote:Quote}) {
    return (
        <Card as="article">
            <CardHeader>
                <Heading as="h1">{quote.sentence}</Heading>
            </CardHeader>
            <CardBody>
                <Text as="p" fontSize="md">
                    {quote.quoteSourceName}
                </Text>
            </CardBody>
        </Card>
    )
}