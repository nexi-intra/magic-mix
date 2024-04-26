"use server";
/*
Parameters

*/
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import { run } from "@/koksmat/magicservices";
import { ShowCodeFragment } from "@/services/ShowCodeFragment";
import { Container Image } from "@/services/magic-mix/models/containerimage";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Container Image
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.containerimage", "search");

  return run<Container Image>(
    "magic-mix.containerimage",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

