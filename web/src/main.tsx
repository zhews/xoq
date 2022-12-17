import {ChakraProvider} from "@chakra-ui/react";
import i18n from "i18next";
import detector from "i18next-browser-languagedetector";
import React from 'react'
import ReactDOM from 'react-dom/client'
import {initReactI18next} from "react-i18next";
import {RouterProvider} from "react-router-dom";
import translationDE from "../locales/de.json";
import translationEN from "../locales/en.json";
import "./main.css";
import router from "./Router";

i18n
    .use(initReactI18next)
    .use(detector)
    .init({
        resources: {
            en: {
                translation: translationEN,
            },
            de: {
                translation: translationDE,
            }
        },
        fallbackLng: "de"
    });

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ChakraProvider>
            <RouterProvider router={router}/>
        </ChakraProvider>
    </React.StrictMode>
)
