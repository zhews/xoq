import {Button, Card, CardBody, CardFooter, CardHeader, Heading, Stack, Text} from "@chakra-ui/react";
import {useTranslation} from "react-i18next";
import {BiStats, BsPlay} from "react-icons/all";
import {Link} from "react-router-dom";

export const Home = () => {
    const {t} = useTranslation();
    return (
        <>
            <Card className={"card"}>
                <CardHeader>
                    <Heading>xoq</Heading>
                </CardHeader>
                <CardBody>
                    <Stack>
                        <Text>{t("home.introduction")}</Text>
                        <Text as={"cite"}>{t("home.question")}</Text>
                        <Text>{t("home.reason")}</Text>
                        <Text>{t("home.how")}</Text>
                        <Text>{t("home.details")}</Text>

                    </Stack>
                </CardBody>
                <CardFooter justify={"space-between"}>
                    <Link to={"/game"}>
                        <Button leftIcon={<BsPlay/>}>{t("button.start")}</Button>
                    </Link>
                    <Link to={"/statistic"}>
                        <Button leftIcon={<BiStats/>}>{t("button.statistic")}</Button>
                    </Link>
                </CardFooter>
            </Card>
        </>
    );
}

export default Home
