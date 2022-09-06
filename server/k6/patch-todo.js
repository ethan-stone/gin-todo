import http from "k6/http";
import { check, sleep } from "k6";
import { encode } from "./jwt.js";

export const options = {
  stages: [
    {
      duration: "5s",
      target: 50
    },
    {
      duration: "10s",
      target: 50
    },
    {
      duration: "5s",
      target: 0
    }
  ],
  thresholds: {
    http_req_duration: ["p(99)<1000", "p(95)<700", "avg<500"]
  }
};

const userId = __ENV.USER_ID;
const jwtSecret = __ENV.SUPABASE_JWT_SECRET;

export function setup() {
  const payload = {
    exp: Math.floor(new Date().getTime() / 1000 + 30 * 60),
    sub: userId
  };
  const jwt = encode(payload, jwtSecret, "HS256");
  const authHeader = `Bearer ${jwt}`;
  const res = http.get(`http://localhost:8080/todo?skip=0&limit=100`, {
    headers: {
      Authorization: authHeader
    }
  });
  return { todosRes: res.json(), authHeader };
}

export default function (data) {
  const todos = data.todosRes.data;
  const todo = todos[Math.floor(Math.random() * todos.length)];
  const res = http.patch(
    `http://localhost:8080/todo/${todo.id}`,
    JSON.stringify({
      description: `Todo: ${Math.round(Math.random() * 1000000)}`
    }),
    {
      headers: {
        "Content-Type": "application/json",
        Authorization: data.authHeader
      }
    }
  );

  check(res, {
    "response code was 200": (res) => res.status === 200
  });

  sleep(2);
}
