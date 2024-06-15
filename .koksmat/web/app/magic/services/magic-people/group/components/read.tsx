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
import { GroupItem } from "../applogic/model";

/* yankiebar */

export default function ReadGroup(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<GroupItem>(
    "magic-people.group",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const group = readResult.data;
  return (
    <div>
      {group && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{group.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{group.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Hidden</div>
            <div>{group.hidden}</div>
          </div>{" "}
          <div>
            <div className="font-bold">NotesId</div>
            <div>{group.notesid}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{group.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{group.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{group.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{group.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{group.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
