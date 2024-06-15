/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { UserForm } from "./form";

import { UserItem } from "../applogic/model";
export default function CreateUser() {
  const user = {} as UserItem;
  return <div>{user && <UserForm user={user} editmode="create" />}</div>;
}
