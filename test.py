from datetime import datetime

import plotly.graph_objs as go

# Sample list of event dictionaries
events = [
    {"time": "08:15:00", "details": {"key1": "value1", "key2": "value2"}},
    {"time": "08:45:00", "details": {"key1": "value3", "key2": "value4"}},
    # ... more events
]

# Sort events by time
events.sort(key=lambda x: x['time'])

# Convert event times to a numerical format (hours + minutes/60) and format details
event_times = [datetime.strptime(event["time"], "%H:%M:%S").hour + datetime.strptime(event["time"], "%H:%M:%S").minute / 60 for event in events]
event_details = ['<br>'.join([f'{key}: {value}' for key, value in event["details"].items()]) for event in events]

# Creating the scatter plot
fig = go.Figure(data=go.Scatter(
    x=event_times,
    y=[1]*len(event_times),  # Simple y-coordinate
    mode='markers',
    text=event_details,
    hoverinfo='text+x',
    marker=dict(size=10, color='blue')
))

# Customizing the layout
fig.update_layout(
    title='Distribution of Events Throughout the Day',
    xaxis=dict(
        title='Time of Day (Hours)',
        tickmode='linear',
        tick0=0,
        dtick=1,  # One tick per hour
        range=[0, 24]  # Covering the full 24 hours
    ),
    yaxis=dict(
        showticklabels=False,
        showgrid=False
    ),
    height=200,  # Narrower figure height
    plot_bgcolor='white'  # White background for better visibility
)

# Show the figure
fig.show()
