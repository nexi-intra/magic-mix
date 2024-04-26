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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Column
 */
export default async function call(transactionId: string ,item: Column) {
  console.log( "magic-mix.column", "update");

  return run<Column>(
    "magic-mix.column",
    ["update" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

