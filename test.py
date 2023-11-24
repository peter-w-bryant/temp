from datetime import datetime

import numpy as np
import plotly.graph_objs as go

# Sample list of event dictionaries
events = [
    {"time": "08:15:00", "details": "Event A"},
    {"time": "08:45:00", "details": "Event B"},
    {"time": "09:30:00", "details": "Event C"},
    {"time": "10:00:00", "details": "Event D"},
    {"time": "10:20:00", "details": "Event E"},
    {"time": "11:50:00", "details": "Event F"},
    # ... more events
]

# Convert event times to decimal hours
event_hours = [datetime.strptime(event["time"], "%H:%M:%S").hour + datetime.strptime(event["time"], "%H:%M:%S").minute / 60 for event in events]
event_details = [event["details"] for event in events]

# Creating the scatter plot
fig = go.Figure(data=go.Scatter(
    x=event_hours,
    y=np.zeros(len(event_hours)),  # All points have the same y-coordinate
    mode='markers',
    text=event_details,
    hoverinfo='text+x',
    marker=dict(
        size=10,
        color='blue'
    )
))

# Customizing the layout
fig.update_layout(
    title='Distribution of Events Throughout the Day',
    xaxis_title='Time of Day (Hours)',
    yaxis=dict(
        showticklabels=False,  # Hide y-axis labels
        showgrid=False,        # Hide y-axis grid
    ),
    xaxis=dict(
        tickmode='array',
        tickvals=list(np.arange(0, 24, 1)),  # Mark every hour
        ticktext=[f'{hour}:00' for hour in range(24)],
        range=[0, 24]  # 24-hour range
    )
)

# Show the figure
fig.show()