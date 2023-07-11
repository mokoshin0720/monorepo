import React from 'react';
import { StyleSheet, Text, TouchableOpacity, ViewStyle } from 'react-native';
import { baseColor, mainColor } from 'src/utils/color';

type ButtonProps = {
    label: string;
    onPress: () => void;
    style?: ViewStyle;
};

export const Button: React.FC<ButtonProps> = ({ label, onPress, style }) => {
    return (
        <TouchableOpacity style={[styles.container, style]} onPress={onPress}>
            <Text style={styles.label}>{label}</Text>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    container: {
        backgroundColor: mainColor,
        borderRadius: 4,
        alignSelf: 'flex-start',
        marginBottom: 24,
    },
    label: {
        fontSize: 16,
        lineHeight: 32,
        paddingVertical: 8,
        paddingHorizontal: 32,
        color: baseColor,
    },
});
