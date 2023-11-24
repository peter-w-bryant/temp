from datetime import datetime

import numpy as np
import plotly.graph_objs as go

# Sample list of event dictionaries
events = [
    {"time": "08:15:00", "details": {"key1": "value1", "key2": "value2"}},
    {"time": "08:45:00", "details": {"key1": "value3", "key2": "value4"}},
    # ... more events
]

# Convert event times to HH:MM format and format details
event_times = [datetime.strptime(event["time"], "%H:%M:%S").strftime('%H:%M') for event in events]
event_details = ['<br>'.join([f'{key}: {value}' for key, value in event["details"].items()]) for event in events]

# Creating the scatter plot
fig = go.Figure(data=go.Scatter(
    x=event_times,
    y=np.zeros(len(event_times)),  # All points have the same y-coordinate
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
    xaxis_title='Time of Day',
    yaxis=dict(
        showticklabels=False,  # Hide y-axis labels
        showgrid=False,        # Hide y-axis grid
    ),
    xaxis=dict(
        tickmode='array',
        tickvals=[f'{hour:02d}:00' for hour in range(24)],  # Mark every hour
        ticktext=[f'{hour:02d}:00' for hour in range(24)],
        range=[event_times[0], event_times[-1]]  # Range based on event times
    )
)

# Show the figure
fig.show()
