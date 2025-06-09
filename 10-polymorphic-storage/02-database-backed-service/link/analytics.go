package link

// The current design is open-ended and allows for adding more
// services without breaking clients. Here's an example.
//
// There could be a service called Analytics that provides
// analytics for the shortened links. It could be used to
// track the number of clicks, the location of the clicks,
// and the devices that the links are clicked on.
//
// Methods:
//   - Stats: Returns statistics for a list of links.
//
// Each returned statistic can include the following information:
//   - Clicks: The number of clicks.
//   - Location: The location of the clicks.
//   - Devices: The devices that the links are clicked on.
//
// The Analytics service could be used to provide insights
// to the clients about the performance of their links.
