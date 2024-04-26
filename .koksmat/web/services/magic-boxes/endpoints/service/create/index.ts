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
import { Service } from "@/services/magic-mix/models/service";
/**
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Service
 */
export default async function call(transactionId: string, item: Service) {
  console.log("magic-mix.service", "create");

  return run<Service>(
    "magic-mix.service",
    ["create", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
