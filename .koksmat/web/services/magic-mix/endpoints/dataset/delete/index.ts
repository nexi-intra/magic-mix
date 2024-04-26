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
 * Delete an existing item
 * 
 * id - The id of the item

 * @returns
 * 
 * Dataset
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.dataset", "delete");

  return run<Dataset>(
    "magic-mix.dataset",
    ["delete" , id],
    transactionId,
    600,
    transactionId
  );
}

