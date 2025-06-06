import React, { useEffect, useState } from "react";

function RelayEventsViewer() {
  const [events, setEvents] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let isMounted = true;
    const fetchEvents = async () => {
      try {
        setIsLoading(true);
        const response = await fetch("http://localhost:8083/events");
        const data = await response.json();
        if (isMounted) setEvents(Array.isArray(data) ? data : []);
      } catch (err) {
        if (isMounted) setError(err.message);
      } finally {
        if (isMounted) setIsLoading(false);
      }
    };

    fetchEvents();
    const interval = setInterval(fetchEvents, 2000);
    return () => {
      isMounted = false;
      clearInterval(interval);
    };
  }, []);

  return (
    <div className="bg-white rounded-lg shadow-md p-6 mb-8">
      <h2 className="text-2xl font-bold mb-4 text-center">Relay Events</h2>
      {isLoading && <p className="text-gray-500">Loading events...</p>}
      {error && <p className="text-red-500">Error: {error}</p>}
      {!isLoading && !error && (
        <div className="overflow-x-auto">
          <table className="min-w-full table-auto border border-gray-400 bg-white">
            <thead className="bg-gray-200">
  <tr>
    <th>Index</th>
    <th>Timestamp</th>
    <th>Source Chain</th>
    <th>Tx Hash</th>
    <th>Amount</th>
  </tr>
</thead>
<tbody>
  {events.length === 0 ? (
    <tr>
      <td colSpan={5} className="text-center text-gray-500">
        No events yet.
      </td>
    </tr>
  ) : (
    events.map((event, idx) => (
      <tr key={(event.TxHash || event.txHash || "") + idx}>
        <td>{event.Index ?? event.index ?? idx + 1}</td>
        <td>
          {event.Timestamp
            ? new Date((event.Timestamp ?? event.timestamp) * 1000).toLocaleString()
            : ""}
        </td>
        <td>{event.SourceChain || event.sourceChain}</td>
        <td style={{ wordBreak: "break-all" }}>{event.TxHash || event.txHash}</td>
        <td>{event.Amount || event.amount}</td>
      </tr>
    ))
  )}
</tbody>
          </table>
        </div>
      )}
    </div>
  );
}

export default RelayEventsViewer;