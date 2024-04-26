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
import { Column } from "@/services/magic-mix/models/column";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Column
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.column", "search");

  return run<Column>(
    "magic-mix.column",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

