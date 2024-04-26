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
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Column
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.column", "read");

  return run<Column>(
    "magic-mix.column",
    ["read" , id],
    transactionId,
    600,
    transactionId
  );
}

