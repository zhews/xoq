import {Button, Card, CardBody, CardFooter, CardHeader, Heading, Link, Text} from "@chakra-ui/react";

import {useTranslation} from "react-i18next";
import {AiOutlineHome} from "react-icons/all";

const NotFound = () => {
    const {t} = useTranslation();
    return (
        <>
            <Card className={"card"}>
                <CardHeader>
                    <Heading>{t("title.notFound")}</Heading>
                </CardHeader>
                <CardBody>
                    <Text>{t("notFound.explanation")}</Text>
                </CardBody>
                <CardFooter>
                    <Link href={"/"}>
                        <Button leftIcon={<AiOutlineHome/>}>{t("button.toHome")}</Button>
                    </Link>
                </CardFooter>
            </Card>

        </>
    )
}

export default NotFound;