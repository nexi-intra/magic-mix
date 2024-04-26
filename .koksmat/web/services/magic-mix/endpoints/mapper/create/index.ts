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
import { Mapper } from "@/services/magic-mix/models/mapper";
/**
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Mapper
 */
export default async function call(transactionId: string ,item: Mapper) {
  console.log( "magic-mix.mapper", "create");

  return run<Mapper>(
    "magic-mix.mapper",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

