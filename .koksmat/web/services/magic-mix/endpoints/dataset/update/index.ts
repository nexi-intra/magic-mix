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
import { Dataset } from "@/services/magic-mix/models/dataset";
/**
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Dataset
 */
export default async function call(transactionId: string ,item: Dataset) {
  console.log( "magic-mix.dataset", "update");

  return run<Dataset>(
    "magic-mix.dataset",
    ["update" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

