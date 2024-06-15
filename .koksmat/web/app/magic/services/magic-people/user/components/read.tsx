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
import { UserItem } from "../applogic/model";

/* yankiebar */

export default function ReadUser(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<UserItem>(
    "magic-people.user",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const user = readResult.data;
  return (
    <div>
      {user && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{user.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{user.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">System</div>
            <div>{user.system_id}</div>
          </div>{" "}
          <div>
            <div className="font-bold">FullName</div>
            <div>{user.fullname}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{user.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{user.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{user.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{user.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{user.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
