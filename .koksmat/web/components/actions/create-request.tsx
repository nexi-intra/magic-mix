"use client";

export interface CreateRequestProps {
  application: string;
  body: { [key: string]: any };
  description: string;
  headers: { [key: string]: any };
  method: string;
  name: string;
  route: string;
  /**
   * Search Index is used for concatenating all searchable fields in a single field making in
   * easier to search
   */
  searchindex: string;
  tenant: string;
}

import React from "react";
import ProcessTransaction from "./process";

export default function CreateRequest(props: {
  transactionid: string;
  request: CreateRequestProps;
}) {
  const { request, transactionid } = props;
  return (
    <div>
      <ProcessTransaction
        payload={request}
        processname="create-request"
        transactionId={transactionid}
      />
    </div>
  );
}
