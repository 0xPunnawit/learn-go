import http from "k6/http";

export let options = {
  vus: 5,
  duration: "5s",
  summaryTrendStats: ["avg", "min", "max", "p(90)"], // ลดข้อมูลส่งเข้า InfluxDB
};

export default function () {
  http.get("http://localhost:8000/hello");
}
