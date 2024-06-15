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
import { PersonItem } from "../applogic/model";

/* yankiebar */

export default function ReadPerson(props: { id: number }) {
  const { id } = props;
  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<PersonItem>(
    "magic-people.person",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const person = readResult.data;
  return (
    <div>
      {person && (
        <div>
          <div>
            <div className="font-bold">Name</div>
            <div>{person.name}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Description</div>
            <div>{person.description}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Email</div>
            <div>{person.email}</div>
          </div>{" "}
          <div>
            <div className="font-bold">DisplayName</div>
            <div>{person.displayname}</div>
          </div>{" "}
          <div>
            <div className="font-bold">FirstName</div>
            <div>{person.firstname}</div>
          </div>{" "}
          <div>
            <div className="font-bold">LastName</div>
            <div>{person.lastname}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Unique Id</div>
            <div>{person.uniqueid}</div>
          </div>{" "}
          <div>
            <div className="font-bold">Nationality</div>
            <div>{person.nationality_id}</div>
          </div>
          <div>
            <div>
              <div className="font-bold">id</div>
              <div>{person.id}</div>
            </div>
            <div>
              <div className="font-bold">created_at</div>
              <div>{person.created_at}</div>
            </div>
            <div>
              <div className="font-bold">created_by</div>
              <div>{person.created_by}</div>
            </div>
            <div>
              <div className="font-bold">updated_at</div>
              <div>{person.updated_at}</div>
            </div>
            <div>
              <div className="font-bold">updated_by</div>
              <div>{person.updated_by}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
