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
import { Transformation } from "@/services/magic-mix/models/transformation";
/**
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Transformation
 */
export default async function call(transactionId: string ,item: Transformation) {
  console.log( "magic-mix.transformation", "create");

  return run<Transformation>(
    "magic-mix.transformation",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

