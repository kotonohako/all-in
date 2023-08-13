import { Card, CardBody, CardHeader, Heading, Text } from "@/app/common";
<<<<<<< HEAD
import { Quote } from "@/generated/buf/kotobako/v1/kotobako_pb";
=======
import { Quote } from "@/app/types";
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64

export default function QuoteDetail({quote}:{quote:Quote}) {
    return (
        <Card as="article">
            <CardHeader>
                <Heading as="h1">{quote.sentence}</Heading>
            </CardHeader>
            <CardBody>
                <Text as="p" fontSize="md">
<<<<<<< HEAD
                    {quote.quoteSourceName}
=======
                    {quote.quote_source_name}
>>>>>>> fa4bc7e6add566c908a877ca03cd93dc08ae8f64
                </Text>
            </CardBody>
        </Card>
    )
}